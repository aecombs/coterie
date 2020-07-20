import React, { useState, useEffect } from 'react';
import User from './User';

import axios from 'axios';

const Profile = (props) => {
  const [userData, setUserData] = useState(null);
  const [errorMessage, setErrorMessage] = useState(null);


  //Updating a user
  const updateUser = (userObj) => {
    axios.put(url, userObj)
    .then((response) => {
      setErrorMessage(`User was updated`);
      window.location.reload();
    })
    
    .catch((error) => {
      setErrorMessage(error.message);
      console.log(`Unable to update user: ${errorMessage}`);
    })
  }

    

  const url = `${process.env.REACT_APP_API_BASE_URL}/users/${props.userID}`



  useEffect(() => {
    axios.get(url)
      .then( (response) => {
        const userDataObj = response.data;
        setUserData(userDataObj);
      })
      .catch((error) => {
        setErrorMessage(error);
        console.log(`There was an error retrieving scriptures: ${errorMessage}`)
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
      submitUserCallback={updateUser}
      />
    } else {
    userComponent = <p className="open-sans">There was an error loading your profile...</p>
  };
  

  return (
    <div className="">
      <div className="d-flex flex-wrap">
        {userComponent}

      </div>
    </div>
  )
}

export default Profile;