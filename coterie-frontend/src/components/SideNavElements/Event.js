import React from 'react';
import PropTypes from 'prop-types';

const Event = (props) => {

  const formatDate = (dateStr) => {
    const formattedDate = new Date(dateStr).toUTCString();
    return formattedDate.substring(0,17)
  }

  return (
    <div className="list-group-item card">
      <div className="card-body">
        <p className="open-sans card-title font-weight-bolder">{props.name}</p>
        <p className="open-sans card-text">{formatDate(props.date)}</p>
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