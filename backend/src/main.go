//Backend Code: Gin Framework
//File: main.go

package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type JobRequest struct {
	Image string `json:"image" binding:"required"`
	GPUs  int    `json:"gpus" binding:"required"`
}

func createKubernetesJob(clientset *kubernetes.Clientset, job JobRequest) error {
	// Define the job spec
	jobName := fmt.Sprintf("job-%s", job.Image)
	k8sJob := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: jobName,
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					RestartPolicy: corev1.RestartPolicyNever,
					Containers: []corev1.Container{
						{
							Name:  jobName,
							Image: job.Image,
							Resources: corev1.ResourceRequirements{
								Limits: corev1.ResourceList{
									"nvidia.com/gpu": *resource.NewQuantity(int64(job.GPUs), resource.DecimalSI),
								},
							},
						},
					},
				},
			},
		},
	}

	// Create the job in Kubernetes
	_, err := clientset.BatchV1().Jobs("default").Create(context.TODO(), k8sJob, metav1.CreateOptions{})
	return err
}

func main() {
	r := gin.Default()

	// Set up Kubernetes client
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Endpoint to submit a job
	r.POST("/submit", func(c *gin.Context) {
		var job JobRequest
		if err := c.ShouldBindJSON(&job); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate GPU count
		if job.GPUs < 1 || job.GPUs > 8 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "GPU count must be between 1 and 8."})
			return
		}

		// Schedule job on Kubernetes
		if err := createKubernetesJob(clientset, job); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Job Submitted Successfully",
			"image":   job.Image,
			"gpus":    job.GPUs,
		})
	})

	// Endpoint to check job status
	r.GET("/status", func(c *gin.Context) {
		jobs, err := clientset.BatchV1().Jobs("default").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		jobStatuses := []gin.H{}
		for _, job := range jobs.Items {
			status := "Unknown"
			if job.Status.Succeeded > 0 {
				status = "Completed"
			} else if job.Status.Active > 0 {
				status = "Running"
			} else if job.Status.Failed > 0 {
				status = "Failed"
			}
			jobStatuses = append(jobStatuses, gin.H{
				"job":    job.Name,
				"status": status,
			})
		}

		c.JSON(http.StatusOK, gin.H{"jobs": jobStatuses})
	})

	r.Run() // Start server on default port 8080
}
