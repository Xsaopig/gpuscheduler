import React from 'react';
import JobForm from './components/JobForm';
import JobStatus from './components/JobStatus';

function App() {
  return (
    <div>
      <h1>Kubernetes GPU Scheduler</h1>
      <JobForm />
      <JobStatus />
    </div>
  );
}

export default App;
