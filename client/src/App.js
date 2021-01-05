import React from 'react';
import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import MainTabs from './Tabs.js';
import Navigation from './Nav.js';

function App() {
  return (
    <div className="App">

      <Navigation />
      <MainTabs />
    </div>
  );
}

export default App;
