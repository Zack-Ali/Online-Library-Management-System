import React, { useState } from "react";
import { useNavigate } from 'react-router-dom';
import axios from "axios";
import toast from "react-hot-toast";
const API_URL = "http://localhost:7788/api"
const Login = () => {
  const [formData, setFormData] = useState({
    email: "",
    password: "",
  });
  const navigate = useNavigate();
  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post(API_URL + "/user/login/", formData);
      console.log(response.data.msg);
      toast.success(response.data.msg);
      // Store the token in local storage for future requests
      localStorage.setItem("token", response.data.token);
      localStorage.setItem("username", response.data.data.username);
      navigate("/")
    } catch (error) {
      console.error("Error logging in", error);
      toast.error(error.response.data.msg);
    }
  };

  return (
    <div className="container">
      <h2>Login</h2>
      <form onSubmit={handleSubmit}>
        <div className="mb-3">
          <label className="form-label">Email address</label>
          <input
            type="email"
            name="email"
            className="form-control"
            value={formData.email}
            onChange={handleChange}
            required
          />
        </div>
        <div className="mb-3">
          <label className="form-label">Password</label>
          <input
            type="password"
            name="password"
            className="form-control"
            value={formData.password}
            onChange={handleChange}
            required
          />
        </div>
        <button type="submit" className="btn btn-primary">
          Login
        </button>
      </form>
    </div>
  );
};

export default Login;
