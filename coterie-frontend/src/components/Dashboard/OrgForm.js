import React, { useState } from 'react';
import PropTypes from 'prop-types';

const OrgForm = (props) => {
  //Form
  const [formFields, setFormFields] = useState({
    "name": 'Enter your organization\'s name',
    "mission_statement": 'Enter you mission statement',
    "total_funds": '0',
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

    props.addOrgCallback(formFields);
    
  };

  return (
    <div className="container">
      <form className="" onSubmit={onFormSubmit}>
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
          {/* <label className="open-sans">Mission Statement</label> */}
          <input
            className="open-sans form-control"
            name="mission_statement"
            type="text"
            value={formFields["mission_statement"]}
            onChange={onMissionChange}
          />
        </div>
        {/* <div className="form-group">
          <input
            className="open-sans form-control"
            name="total_funds"
            type="text"
            value={formFields["total_funds"]}
            onChange={onFundsChange}
          />
        </div> */}
        <input 
          className="btn btn-light w-100 text-center" 
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