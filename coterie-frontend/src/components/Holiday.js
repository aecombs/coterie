import React from 'react';
import PropTypes from 'prop-types';

const Holiday = (props) => {
  return (
    <div className="list-group-item card">
      <div className="card-body">
        <p className="open-sans card-title">{props.name}</p>
        <p className="open-sans card-text">{props.date}</p>
      </div>
      <div className="card-body">
        <p className="open-sans card-text hidden-until-hover">{props.description}</p>
      </div>
    </div>
  )
}

Holiday.propTypes = {
  id: PropTypes.number,
  name: PropTypes.string,
  date: PropTypes.string,
  description: PropTypes.string,
  orgID: PropTypes.number
};

export default Holiday;