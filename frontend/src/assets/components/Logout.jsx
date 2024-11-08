import React from 'react';
import { Link } from 'react-router-dom';

const Logout = () => {
  return (
    <div>
      <h2>Logout Page</h2>
      <p>You have been logged out successfully.</p>
      <Link to={"/login"}>Login page</Link>
    </div>
  );
};

export default Logout;
