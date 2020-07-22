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
          <div className="row justify-content-between my-2 card border-0">
          <button className={visibility ? "hidden" : "btn btn-outline-light text-dark open-sans mb-2"} onClick={toggleFormVisibility}>
            <div className="col card-text pt-2">
              <h4 className={visibility ? "hidden" : "mt-3 text-center open-sans"}>{props.name}</h4>
              <div className="container w-75 mt-4 mb-5">
                 <small className="open-sans text-muted text-left">Full Bio</small>
                <p className={visibility ? "hidden" : "open-sans"}>{props.bio}</p>
              </div>
            
            </div>
          </button>
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