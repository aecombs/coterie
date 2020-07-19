import React, { useState } from 'react';
import PropTypes from 'prop-types';

const OrgForm = (props) => {
  //Form
  const [formFields, setFormFields] = useState({
    "name": '',
    "mission_statement": '',
    "total_funds": '',
  });

  const onNameChange = (event) => {
    console.log(`Text field updated ${event.target.value}`);
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

  const onFundsChange = (event) => {
    console.log(`total_funds field updated ${event.target.value}`);
    setFormFields({
      ...formFields,
      "total_funds": event.target.value,
    });
  };
  
//callback func to an http req for posting org.
  const onFormSubmit = (event) => {
    event.preventDefault();

    props.addOrgCallback(formFields);
    
  };

  return (
    <div className="container">
      <form className="" onSubmit={onFormSubmit}>
        <h2 className="">New Organization</h2>
        <div className="form-group">
            <label className="open-sans">Name</label>
            <input
              className="open-sans form-control"
              name="name"
              type="text"
              value={formFields["name"]}
              onChange={onNameChange}
            />
        </div>
        <div className="form-group">
          <label className="open-sans">Mission Statement</label>
          <input
            className="open-sans form-control"
            name="mission_statement"
            type="text"
            value={formFields["mission_statement"]}
            onChange={onMissionChange}
          />
        </div>
        <div className="form-group">
          <label className="open-sans">Total Funds</label>
          <input
            className="open-sans form-control"
            name="total_funds"
            type="text"
            value={formFields["total_funds"]}
            onChange={onFundsChange}
          />
        </div>
        <input 
          className="btn btn-light" 
          type="submit"
          value="Add Organization"
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