import React from 'react';
import PropTypes from 'prop-types';

const Event = (props) => {
  return (
    <div className="list-group-item card">
      <div className="card-body">
        <p className="open-sans card-title">{props.name}</p>
        <p className="open-sans card-text">{props.date}</p>
        <p className="open-sans card-text">{props.occasion}</p>
      </div>
      <div className="card-body">
        <p className="open-sans card-text hidden-until-hover">{props.description}</p>
      </div>
    </div>
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