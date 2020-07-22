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
    <section className="w-100 py-2 container mr-2 d-flex justify-content-start">
      <div className="w-100">
        <img src={props.avatar} alt="Avatar" className="profile-image row img-thumbnail" />
        <div className="row justify-content-between my-2 card">
          <div className="col card-text pt-2">
            <p className={visibility ? "hidden" : "open-sans"}>{props.name}</p>
            <p className={visibility ? "hidden" : "open-sans"}>{props.bio}</p>
          </div>
          <div className="col card-text text-right">
            <button className={visibility ? "hidden" : "btn btn-secondary open-sans mb-2"} onClick={toggleFormVisibility}>Update</button>
          </div>
          <UpdateProfile 
            visibility={visibility}
            submitUserCallback={props.submitUserCallback}
            onSubmitCallback={toggleFormVisibility}
            name={props.name}
            email={props.email}
            bio={props.bio}
          />
        </div>        
      </div>
    </section>
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