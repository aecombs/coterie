import React, { useState } from 'react';
import PropTypes from 'prop-types';

const NewMember = (props) => {
  //Form
  const [formFields, setFormFields] = useState({
    "name": '',
    "birthdate": "1999-01-01",
    "class": 'Embraced',
    "email": '',
    "funds_raised": '',
    "organization_id": props.orgID,
  });

  const onNameChange = (event) => {
    console.log(`name field updated ${event.target.value}`);
    setFormFields({
      ...formFields,
      "name": event.target.value,
    });
  };

  const onBirthdateChange = (event) => {
    console.log(`birthdate field updated ${event.target.value}`);
    setFormFields({
      ...formFields,
      "birthdate": event.target.value,
    });
  };
  const onClassChange = (event) => {
    console.log(`class field updated ${event.target.value}`);
    setFormFields({
      ...formFields,
      "class": event.target.value,
    });
  };
  const onEmailChange = (event) => {
    console.log(`email field updated ${event.target.value}`);
    setFormFields({
      ...formFields,
      "email": event.target.value,
    });
  };
  const onFundsChange = (event) => {
    console.log(`funds_raised field updated ${event.target.value}`);
    setFormFields({
      ...formFields,
      "funds_raised": event.target.value.substring(1),
    });
  };
  
//callback func to an http req for posting org.
  const onFormSubmit = (event) => {
    event.preventDefault();
    props.onSubmitCallback();

    const beforeFunds = formFields["funds_raised"].substr(0);
    formFields["funds_raised"] = beforeFunds.replace(/[^\d]/gi, '');

    props.addMemberCallback(formFields);
    
  };

  return (
    <div className={ props.visibility ? "container py-3" : "container hidden"}>
      <form className="" onSubmit={onFormSubmit}>
        <div className="form-group">
          <small className="open-sans form-text text-muted">Member's Name</small>
          <input
            className="open-sans form-control"
            name="name"
            type="text"
            placeholder="Joe Shmoe"
            value={formFields["name"]}
            onChange={onNameChange}
          />
        </div>
        <div className="form-group">
          <small className="open-sans form-text text-muted">Member's Birthday</small>
          <input
            className="open-sans form-control"
            name="birthdate"
            type="date"
            value={formFields["birthdate"]}
            onChange={onBirthdateChange}
          />
        </div>
        <div className="form-group">
          <small className="open-sans form-text text-muted">Member's Class</small>
          <select className="open-sans form-control" 
            name="class"
            value={formFields["class"]}
            onChange={onClassChange}
          >
            <option>Revered</option>
            <option>Respected</option>
            <option>Embraced</option>
            <option>Tolerated</option>
            <option>Shamed</option>
          </select>
        </div>
        <div className="form-group">
          <small className="open-sans form-text text-muted">Member's Email</small>
          <input
            className="open-sans form-control"
            name="email"
            type="email"
            placeholder="joe@shmoe.com"
            value={formFields["email"]}
            onChange={onEmailChange}
          />
        </div>
        <div className="form-group">
        <small className="open-sans form-text text-muted">Funds raised by this member</small>
          <input
            className="open-sans form-control"
            name="funds_raised"
            type="text"
            placeholder="$50.00"
            value={"$"+formFields["funds_raised"]}
            onChange={onFundsChange}
          />
        </div>
        <input 
          className="btn btn-light w-100 text-center" 
          type="submit"
          value="Add Member"
          onClick={onFormSubmit}
        />
      </form>
    </div>
)
}

NewMember.propTypes = {
  orgID: PropTypes.number,
  visibility: PropTypes.bool,
};

export default NewMember;