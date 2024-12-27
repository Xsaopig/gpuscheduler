import React, { useEffect, useState } from 'react';

function JobStatus() {
  const [jobs, setJobs] = useState([]);

  useEffect(() => {
    const fetchStatus = async () => {
      const response = await fetch('/status');
      const data = await response.json();
      setJobs(data.jobs);
    };
    fetchStatus();
  }, []);

  return (
    <div>
      <h2>Job Status</h2>
      <ul>
        {jobs.map((job, index) => (
          <li key={index}>
            {job.job}: {job.status}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default JobStatus;
