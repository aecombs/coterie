import React, { useState } from 'react';
import PropTypes from 'prop-types';

const ChapterForm = (props) => {
  //Form
  const [formFields, setFormFields] = useState({
    "id": (props.id ? props.id : null),
    "name": (props.name ? props.name : ""),
    "text": (props.text ? props.text : ""),
    "position": (props.position ? props.position : (props.scripLength + 1)),
    "scripture_id": props.scripID,
  });

  const onNameChange = (event) => {
    console.log(`name field updated ${event.target.value}`);
    setFormFields({
      ...formFields,
      "name": event.target.value,
    });
  };

  const onTextChange = (event) => {
    console.log(`text field updated ${event.target.value}`);
    setFormFields({
      ...formFields,
      "text": event.target.value,
    });
  };

  // const onPositionChange = (event) => {
  //   console.log(`position field updated ${event.target.value}`);
  //   setFormFields({
  //     ...formFields,
  //     "position": event.target.value,
  //   });
  // };

//callback func to an http req for posting org.
  const onFormSubmit = (event) => {
    event.preventDefault();
    props.onSubmitCallback();

    props.submitChapterCallback(formFields);
    
  };

  return (
    <div className={ props.visibility ? "py-3" : "hidden"}>
      <form className="" onSubmit={""}>
        <div className="form-group">
          <small className="open-sans form-text text-muted">Chapter Title</small>
          <input
            className="open-sans form-control"
            name="name"
            type="text"
            placeholder="Type in the chapter's title"
            value={formFields["name"]}
            onChange={onNameChange}
          />
        </div>
        <div className="form-group">
          <small className="open-sans form-text text-muted">Chapter Text</small>
          <textarea
            className="open-sans form-control"
            name="text"
            placeholder="Type in the chapter's text here"
            rows="5"
            value={formFields["text"]}
            onChange={onTextChange}
          />
        </div>
        {/* <div className="form-group">
          <small className="open-sans form-text text-muted">Position in Scripture</small>
          <input
            className="open-sans form-control"
            name="date"
            type="date"
            value={formFields["date"]}
            onChange={onPositionChange}
          />
        </div> */}
        <div className="btn-group w-100">
          <input 
            className="btn btn-success text-center" 
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

ChapterForm.propTypes = {
  id: PropTypes.number,
  name: PropTypes.string,
  text: PropTypes.string,
  position: PropTypes.number,
  scripLength: PropTypes.number,
  scripID: PropTypes.number,
  visibility: PropTypes.bool,
};

export default ChapterForm;