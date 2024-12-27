import React, { useState } from 'react';

function JobForm() {
  const [image, setImage] = useState('');
  const [gpus, setGpus] = useState(1);

  const handleSubmit = async (e) => {
    e.preventDefault();
    const payload = { image, gpus: parseInt(gpus, 10) }; // 构建请求数据
    try {
      // 发送 POST 请求到后端
      const response = await axios.post(`${process.env.REACT_APP_BACKEND_URL}/submit`, payload);
        alert(response.data.message); // 显示后端返回的消息
      } catch (error) {
        console.error("Error submitting job:", error); // 打印错误日志
        alert("Failed to submit job: " + error.response?.data?.error || error.message);
      }
    };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label>Docker Image:</label>
        <input
          type="text"
          value={image}
          onChange={(e) => setImage(e.target.value)}
          required
        />
      </div>
      <div>
        <label>GPU Count:</label>
        <input
          type="number"
          value={gpus}
          onChange={(e) => setGpus(parseInt(e.target.value, 10))}
          min="1"
          required
        />
      </div>
      <button type="submit">Submit</button>
    </form>
  );
}

export default JobForm;
