import React from 'react';
import PropTypes from 'prop-types';

const Event = (props) => {
  const formatDate = (dateStr) => {
    const formattedDate = new Date(dateStr).toUTCString();
    return formattedDate.substring(0,17)
  }

  return (
    <div className="list-group-item card">
      <div className="card-body row justify-content-between">
        <div className="">
          <p className="card-title font-weight-bolder">{props.name}</p>
          <p className="open-sans card-text">{props.description}</p>
        </div>
        <div>
          <p className="open-sans card-text">{formatDate(props.date)}</p>
        <p className="open-sans card-text">{props.occasion}</p>
        </div>
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