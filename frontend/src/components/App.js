import React from "react";
import Header from "./Header";
import { Navigate, Route, Routes } from "react-router-dom";
import Login from "./Login";
import Dating from "./Dating";
import Details from "./Details";
import Profile from "./Profile";
import { AUTH_TOKEN } from "../constants";

const ProtectedRoute = ({ children }) => {
  const token = localStorage.getItem(AUTH_TOKEN);

  if (!token) {
    return <Navigate to="/" replace />;
  }

  return children;
};

const App = () => (
  <div className="center w85">
    <Header />
    <div className="ph3 pv1 background-gray">
      <Routes>
        <Route path="/" element={<Navigate replace to="/login" />} />
        <Route path="/login" element={<Login />} />
        <Route
          path="/dating"
          element={
            <ProtectedRoute>
              <Dating />
            </ProtectedRoute>
          }
        />

        <Route
          path="/details"
          element={
            <ProtectedRoute>
              <Details />
            </ProtectedRoute>
          }
        />
        <Route
          path="/profile"
          element={
            <ProtectedRoute>
              <Profile />
            </ProtectedRoute>
          }
        />
      </Routes>
    </div>
  </div>
);

export default App;
