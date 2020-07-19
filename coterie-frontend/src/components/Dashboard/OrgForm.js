import React, { useState, useEffect } from 'react';
import PropTypes from 'prop-types';

const OrgForm = (props) => {
  //Form
  const [formFields, setFormFields] = useState({
    name: '',
    mission_statement: '',
    total_funds: '',
  });

  const onNameChange = (event) => {
    console.log(`Text field updated ${event.target.value}`);
    setFormFields({
      ...formFields,
      text: event.target.value,
    });
  };

  const onMissionChange = (event) => {
    console.log(`mission_statement field updated ${event.target.value}`);
    setFormFields({
      ...formFields,
      mission_statement: event.target.value,
    });
  };

  const onFundsChange = (event) => {
    console.log(`total_funds field updated ${event.target.value}`);
    setFormFields({
      ...formFields,
      total_funds: event.target.value,
    });
  };


  // const dataChanged = (data) => {
  //   // data = { description: "New validated text comes here" }
  //   // Update your model from here
  //   console.log(data)
  //   this.setState({...data})
  // }
  
//callback func to an http req for posting org.
  const onFormSubmit = (event) => {
    event.preventDefault();

    props.addOrgCallback(formFields);
    
    setFormFields({
      name: '',
      mission_statement: '',
      total_funds: '',
    });
  };

  return (
    <div className="container">
      <form className="" onSubmit={onFormSubmit}>
        <h2 className="">New Organization</h2>
        <div>
            <label className="">Name</label>
            <input
              className=""
              name="name"
              type="text"
              value={formFields.name}
              onChange={onNameChange}
            />
        </div>
        <div>
          <label className="">Mission Statement</label>
          <input
            className=""
            name="mission_statement"
            type="text"
            value={formFields.name}
            onChange={onMissionChange}
          />
        </div>
        <div>
          <label className="">Total Funds</label>
          <input
            className=""
            name="total_funds"
            type="text"
            value={formFields.name}
            onChange={onFundsChange}
          />
        </div>
        <input 
          className="btn btn-dark" 
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