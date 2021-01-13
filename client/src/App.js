import React from 'react';
import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import Config from './components/Config.js';
import Navigation from './components/Nav.js';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import { CurrentTriage } from './components/Triage';

function App() {
  return (
    <div className="App">

      <Navigation />
      <BrowserRouter>
      <Switch>
          <Route path="/config">
            <Config />
          </Route>
          <Route path="/triage">
            <CurrentTriage />
          </Route>
          <Route path="/whale">
          </Route>
        </Switch>
      </BrowserRouter>
      
    </div>
  );
}

export default App;
