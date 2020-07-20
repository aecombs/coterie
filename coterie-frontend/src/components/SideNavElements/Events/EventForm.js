import React, { useState } from 'react';
import PropTypes from 'prop-types';

const EventForm = (props) => {
  //Form
  const [formFields, setFormFields] = useState({
    "id": (props.id ? props.id : null),
    "name": (props.name ? props.name : ""),
    "occasion": (props.occasion ? props.occasion : "Picnic"),
    "date": (props.date ? props.date : ""),
    "description": (props.description ? props.description : ""),
    "organization_id": props.orgID,
  });

  const onNameChange = (event) => {
    console.log(`name field updated ${event.target.value}`);
    setFormFields({
      ...formFields,
      "name": event.target.value,
    });
  };

 
  const onOccasionChange = (event) => {
    console.log(`occasion field updated ${event.target.value}`);
    setFormFields({
      ...formFields,
      "occasion": event.target.value,
    });
  };

  const onDateChange = (event) => {
    console.log(`date field updated ${event.target.value}`);
    setFormFields({
      ...formFields,
      "date": event.target.value,
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

    props.submitEventCallback(formFields);

    
  };

  return (
    <div className={ props.visibility ? "py-3" : "hidden"}>
      <form className="" onSubmit={""}>
        <div className="form-group">
          <small className="open-sans form-text text-muted">Event Name</small>
          <input
            className="open-sans form-control"
            name="name"
            type="text"
            placeholder="Give your event a name"
            value={formFields["name"]}
            onChange={onNameChange}
          />
        </div>
        <div className="form-group">
          <small className="open-sans form-text text-muted">Event Occasion</small>
          <select className="open-sans form-control" 
            name="occasion"
            value={formFields["occasion"]}
            onChange={onOccasionChange}
          >
            <option>Picnic</option>
            <option>Ritual</option>
            <option>Collective Meditation</option>
            <option>Holiday</option>
            <option>Meeting</option>
            <option>Emergency Meeting</option>
          </select>
        </div>
        <div className="form-group">
          <small className="open-sans form-text text-muted">Event Date</small>
          <input
            className="open-sans form-control"
            name="date"
            type="date"
            value={formFields["date"]}
            onChange={onDateChange}
          />
        </div>
        <div className="form-group">
          <small className="open-sans form-text text-muted">Event Description</small>
          <textarea
            className="open-sans form-control"
            name="description"
            placeholder="What should members expect at this event?"
            rows="5"
            value={formFields["description"]}
            onChange={onDescriptionChange}
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

EventForm.propTypes = {
  id: PropTypes.number,
  name: PropTypes.string,
  occasion: PropTypes.string,
  date: PropTypes.string,
  description: PropTypes.string,
  orgID: PropTypes.number,
  visibility: PropTypes.bool,
};

export default EventForm;