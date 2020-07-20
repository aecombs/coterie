import React, { useState } from 'react';
import PropTypes from 'prop-types';
import EventForm from './EventForm';

const Event = (props) => {
  const [updateMode, setUpdateMode] = useState(false);

  const formatDate = (dateStr) => {
    const formattedDate = new Date(dateStr).toUTCString();
    return formattedDate.substring(0,17)
  }

  const updateEvent = (eventObj) => {
    props.submitEventCallback(eventObj)
  }

  const setUpdate = () => {
    setUpdateMode(!updateMode);
  }

  return (
    <section>
      <button onClick={setUpdate} className={ updateMode ? "hidden" : "btn list-group-item list-group-item-action"}>
        <div className="card-body row justify-content-between">
          <div>
            <h5 className="card-title font-weight-bolder text-left">{props.name}</h5>
            <p className="open-sans card-text text-left">{props.description}</p>
          </div>
          <div>
            <p className="open-sans card-text text-right">{formatDate(props.date)}</p>
          <p className="open-sans card-text text-right">{props.occasion}</p>
          </div>
        </div>
      </button>
      <EventForm 
        id={props.id}
        name={props.name}
        occasion={props.occasion}
        date={props.date}
        description={props.description}
        orgID={props.orgID}
        visibility={updateMode}
        submitEventCallback={updateEvent}
        onSubmitCallback={setUpdate}
        />
    </section>
  )
}

Event.propTypes = {
  id: PropTypes.number,
  name: PropTypes.string,
  occasion: PropTypes.string,
  date: PropTypes.string,
  description: PropTypes.string,
  orgID: PropTypes.number
};

export default Event;