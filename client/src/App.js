import React from 'react';
import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';

import Navigation from './components/Nav.js';
import MainTabs from './components/Tabs';


function App() {
  return (
    <div className="App">

      <Navigation />
      <MainTabs />
      
    </div>
  );
}

export default App;
