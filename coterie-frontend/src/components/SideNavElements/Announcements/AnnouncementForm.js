import React, { useState } from 'react';
import PropTypes from 'prop-types';

const AnnouncementForm = (props) => {
  //Form
  const [formFields, setFormFields] = useState({
    "id": (props.id ? props.id : null),
    "header": (props.header ? props.header : ""),
    "description": (props.description ? props.description : ""),
    "date": (props.date ? props.date : new Date(Date.now())),
    "organization_id": props.orgID,
  });

  const onHeaderChange = (event) => {
    console.log(`header field updated ${event.target.value}`);
    setFormFields({
      ...formFields,
      "header": event.target.value,
    });
  };

  const onDescriptionChange = (event) => {
    console.log(`description field updated ${event.target.value}`);
    setFormFields({
      ...formFields,
      "description": event.target.value,
    });
  };

//callback func to an http req for posting org.
  const onFormSubmit = (event) => {
    event.preventDefault();
    props.onSubmitCallback();

    props.submitAnnouncementCallback(formFields);

    
  };

  return (
    <div className={ props.visibility ? "py-3" : "hidden"}>
      <form className="" onSubmit={""}>
        <div className="form-group">
          <small className="open-sans form-text text-muted">Announcement Name</small>
          <input
            className="open-sans form-control"
            name="header"
            type="text"
            placeholder="What header should this announcement have?"
            value={formFields["header"]}
            onChange={onHeaderChange}
          />
        </div>
        <div className="form-group">
          <small className="open-sans form-text text-muted">Announcement Description</small>
          <textarea
            className="open-sans form-control"
            name="description"
            placeholder="What would you like to announce?"
            rows="5"
            value={formFields["description"]}
            onChange={onDescriptionChange}
          />
        </div>
        <div className="btn-group w-100">
          <input 
            className="btn btn-success text-center w-75" 
            type="submit"
            value="Save"
            onClick={onFormSubmit}
          />
          <input 
            className="btn btn-light text-center w-25" 
            type="submit"
            value="Cancel"
            onClick={props.onSubmitCallback}
          />
          <input 
            className={ props.id ? "btn btn-danger delete-button float-right" : "hidden"} 
            type="submit"
            value="Delete"
            onClick={props.deleteAnnouncementCallback}
          />
        </div>
      </form>
    </div>
)
}

AnnouncementForm.propTypes = {
  id: PropTypes.number,
  header: PropTypes.string,
  description: PropTypes.string,
  date: PropTypes.string,
  orgID: PropTypes.number,
  visibility: PropTypes.bool,
};

export default AnnouncementForm;