import React, { useState } from 'react';
import UpdateProfile from './UpdateProfile';
import PropTypes from 'prop-types';

const User = (props) => {
  const [visibility, setVisibility] = useState(false);
    //toggle visibility of the updateUser component
  const toggleFormVisibility = () => {
    setVisibility(!visibility);
    return;
  }

    
  return (
    <div className="w-100 py-2">
      <img src={props.avatar} alt="Avatar" className="w-25" />
      <div className="column">
        <h5 className={visibility ? "hidden" : "py-3 open-sans"}>{props.name}</h5>

        <UpdateProfile 
          visibility={visibility}
          updateUserCallback={props.updateUserCallback}
          onSubmitCallback={toggleFormVisibility}
          name={props.name}
          email={props.email}
          bio={props.bio}
        />
        <p className={visibility ? "hidden" : "open-sans"}>{props.bio}</p>
      </div>
      <div className="column">
        <button className={visibility ? "hidden" : "btn btn-secondary open-sans"} onClick={toggleFormVisibility}>Update Profile</button> 
      </div>
      
    </div>
  )
}

User.propTypes = {
  id: PropTypes.number,
  name: PropTypes.string,
  email: PropTypes.string,
  bio: PropTypes.string,
  avatar: PropTypes.string
};

export default User;