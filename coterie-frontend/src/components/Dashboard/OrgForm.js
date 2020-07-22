import React, { useState } from 'react';
import PropTypes from 'prop-types';

const OrgForm = (props) => {
  //Form
  const [formFields, setFormFields] = useState({
    "id": (props.id ? props.id : null),
    "name": (props.name ? props.name : 'Your organization\'s name here'),
    "mission_statement": (props.missionStatement ? props.missionStatement : ''),
    "user_id": props.userID,
  });

  const onNameChange = (event) => {
    console.log(`name field updated ${event.target.value}`);
    setFormFields({
      ...formFields,
      "name": event.target.value,
    });
  };

  const onMissionChange = (event) => {
    console.log(`mission_statement field updated ${event.target.value}`);
    setFormFields({
      ...formFields,
      "mission_statement": event.target.value,
    });
  };
  
//callback func to an http req for posting org.
  const onFormSubmit = (event) => {
    event.preventDefault();
    props.onSubmitCallback();

    props.submitOrgCallback(formFields);
    
  };

  return (
    <div className={props.visibility ? "container row" : "hidden"}>
      <form className="w-100" onSubmit={""}>
        <h2 className="form-title">{formFields["name"]}</h2>
        <div className="form-group">
            {/* <label className="open-sans">Name</label> */}
            <input
              className="open-sans form-control"
              name="name"
              type="text"
              value={formFields["name"]}
              onChange={onNameChange}
            />
        </div>
        <div className="form-group">
          <input
            className="open-sans form-control"
            name="mission_statement"
            type="text"
            placeholder="Enter you mission statement"
            value={formFields["mission_statement"]}
            onChange={onMissionChange}
          />
        </div>
        <input 
          className="btn btn-light w-100 text-center" 
          type="submit"
          value="Submit"
          onClick={onFormSubmit}
        />
      </form>
    </div>
)
}

OrgForm.propTypes = {
  selection: PropTypes.string,
  userID: PropTypes.number,
};

export default OrgForm;