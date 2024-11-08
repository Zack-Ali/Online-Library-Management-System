import React, { useState, useEffect } from "react";
import axios from "axios";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import toast from "react-hot-toast";
import Profile from "./assets/components/Profile";
import Logout from "./assets/components/Logout";
import Signup from "./assets/auth/Signup";
import Login from "./assets/auth/Login";
import Header from "./assets/components/Header";
import HomePage from "./assets/pages/HomePage";
import BookPage from "./assets/pages/BookPage";

export const API_URL = "http://localhost:7788/api";

const App = () => {  
  return (
    <Router>
      {/* Header Component */}
      <Header />

      {/* Routes for different pages */}
      <div className="container mt-5">
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/book/add" element={<BookPage />} />
          <Route path="/book/edit/:id" element={<BookPage />} />
          <Route path="/profile" element={<Profile />} />
          <Route path="/logout" element={<Logout />} />
          <Route path="/signup" element={<Signup />} /> {/* Signup route */}
          <Route path="/login" element={<Login />} /> {/* Login route */}
        </Routes>
      </div>
    </Router>
  );
};

export default App;
