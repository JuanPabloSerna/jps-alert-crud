import React from 'react'
import {
  BrowserRouter as Router,
  Routes,
  Route,
} from "react-router-dom";
import List from './alerts/List';
import ViewAlert from './alerts/ViewAlert'; 

export default function App() {

  return (
    <Router>
      <Routes>
        <Route index element={<List />} />
        <Route path="/alert/:id" element={<ViewAlert />} />
      </Routes>
    </Router>
  );
}

