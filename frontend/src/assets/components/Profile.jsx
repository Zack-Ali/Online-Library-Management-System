import React from 'react';

const Profile = () => {
  const username = localStorage.getItem("username");
  return (
    <div>
      <h2>Profile Page</h2>
      <p>Welcome <strong>{username} </strong>to your profile. Here you can update your personal information.</p>
    </div>
  );
};

export default Profile;
