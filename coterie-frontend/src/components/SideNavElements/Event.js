import React from 'react';
import PropTypes from 'prop-types';

const Event = (props) => {
  const formatDate = (dateStr) => {
    const formattedDate = new Date(dateStr).toUTCString();
    return formattedDate.substring(0,17)
  }

  return (
    <button className="btn list-group-item list-group-item-action">
      <div className="card-body row justify-content-between">
        <div>
          <p className="card-title font-weight-bolder text-left">{props.name}</p>
          <p className="open-sans card-text text-left">{props.description}</p>
        </div>
        <div>
          <p className="open-sans card-text text-right">{formatDate(props.date)}</p>
        <p className="open-sans card-text text-right">{props.occasion}</p>
        </div>
      </div>
    </button>
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