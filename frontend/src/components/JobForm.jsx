import React, { useState } from 'react';

function JobForm() {
  const [image, setImage] = useState('');
  const [gpus, setGpus] = useState(1);

  const handleSubmit = async (e) => {
    e.preventDefault();
    const response = await fetch('/submit', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ image, gpus }),
    });
    const data = await response.json();
    alert(data.message);
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
