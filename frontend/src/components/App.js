import React from 'react';
import Header from './Header';
import {Navigate, Route, Routes} from 'react-router-dom';
import Login from './Login';
import Dating from './Dating';
import Details from './Details';
import Profile from './Profile';

const App = () => (
  <div className="center w85">
    <Header />
    <div className="ph3 pv1 background-gray">
      <Routes>
        <Route
          path="/"
          element={<Navigate replace to="/login" />}
        />
        <Route path="/login" element={<Login/>} />
        <Route path="/dating" element={<Dating/>} />
        <Route path="/details" element={<Details/>} />
        <Route path="/profile" element={<Profile/>} />
      </Routes>
    </div>
  </div>
);

export default App;
