import React, { useState } from 'react';
import PropTypes from 'prop-types';

const ScriptureForm = (props) => {
  //Form
  const [formFields, setFormFields] = useState({
    "id": (props.id ? props.id : null),
    "name": (props.name ? props.name : ""),
    "organization_id": props.orgID,
  });

  const onNameChange = (event) => {
    console.log(`name field updated ${event.target.value}`);
    setFormFields({
      ...formFields,
      "name": event.target.value,
    });
  };

//callback func to an http req for posting org.
  const onFormSubmit = (event) => {
    event.preventDefault();
    props.onSubmitCallback();

    if (props.id !== null) {
      props.updateScriptureCallback(formFields);
    } else {
      props.addScriptureCallback(formFields);
    }
    
  };

  return (
    <div className={ props.visibility ? "py-3" : "hidden"}>
      <form className="" onSubmit={""}>
        <div className="form-group">
          <small className="open-sans form-text text-muted">Scripture Title</small>
          <input
            className="open-sans form-control"
            name="name"
            type="text"
            placeholder="Name your new scripture"
            value={formFields["name"]}
            onChange={onNameChange}
          />
        </div>
        <div className="btn-group w-100">
          <input 
            className="btn btn-success text-center w-50" 
            type="submit"
            value="Save"
            onClick={onFormSubmit}
          />
          <input 
            className="btn btn-light text-center" 
            type="submit"
            value="Cancel"
            onClick={props.onSubmitCallback}
          />
        </div>
      </form>
    </div>
)
}

ScriptureForm.propTypes = {
  id: PropTypes.number,
  name: PropTypes.string,
  orgID: PropTypes.number,
  visibility: PropTypes.bool,
};

export default ScriptureForm;