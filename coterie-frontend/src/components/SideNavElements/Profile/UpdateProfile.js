import React, { useState } from 'react';
import PropTypes from 'prop-types';

const UpdateProfile = (props) => {
  //Form
  const [formFields, setFormFields] = useState({
    "name": props.name,
    "email": props.email,
    "bio": props.bio,
  });

  const onNameChange = (event) => {
    console.log(`name field updated ${event.target.value}`);
    setFormFields({
      ...formFields,
      "name": event.target.value,
    });
  };

  const onBioChange = (event) => {
    console.log(`bio field updated ${event.target.value}`);
    setFormFields({
      ...formFields,
      "bio": event.target.value,
    });
  };
  
//callback func to an http req for posting org.
  const onFormSubmit = (event) => {
    event.preventDefault();
    props.onSubmitCallback();

    props.updateUserCallback(formFields);
    
  };

  return (
    <div className={ props.visibility ? "container mt-n2" : "hidden"}>
      <form className="" onSubmit={""}>
        <div className="form-group">
          {/* <small className="open-sans form-text text-muted">Name</small> */}
          <input
            className="open-sans form-control"
            name="name"
            type="text"
            value={formFields["name"]}
            onChange={onNameChange}
          />
        </div>
        <div className="form-group">
          {/* <small className="open-sans form-text text-muted">Bio</small> */}
          <textarea
            className="open-sans form-control"
            name="bio"
            rows="5"
            value={formFields["bio"]}
            onChange={onBioChange}
          />
        </div>
        <input 
          className="btn btn-light w-100 text-center" 
          type="submit"
          value="Update"
          onClick={onFormSubmit}
        />
      </form>
    </div>
)
}

UpdateProfile.propTypes = {
  userID: PropTypes.number,
};

export default UpdateProfile;