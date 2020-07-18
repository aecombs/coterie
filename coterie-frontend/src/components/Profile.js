import React, { useState, useEffect } from 'react';
import User from './User';
import axios from 'axios';

const Profile = (props) => {
  const [userData, setUserData] = useState(null);

  const url = `http://localhost:3000/users/${props.userID}`

  useEffect(() => {
    axios.get(url)
      .then( (response) => {
        const userDataObj = response.data;
        setUserData(userDataObj);
      })
      .catch((error) => {
        console.log(`There was an error retrieving scriptures: ${error}`)
      });
  },[url])
  
  let userComponent = undefined
  if (userData !== null && userData["id"] == props.userID) {
    userComponent = 
      <User 
      key={userData["id"]}
      id={userData["id"]}
      name={userData["name"]}
      email={userData["email"]}
      bio={userData["bio"]}
      avatar={userData["avatar"]}
      />
    } else {
    userComponent = <p className="open-sans">There was an error loading your profile...</p>
  };

  return (
    <div className="container">
      <h4 className=" w-100">Profile</h4>
      <div className="d-flex">{userComponent}</div>
      
    </div>
  )
}

export default Profile;