import React from 'react';
import PropTypes from 'prop-types';

const Announcement = (props) => {
  return (
    <div className="list-group-item card">
      <div className="card-body">
        <p className="open-sans card-title">{props.header}</p>
      
        <p className="open-sans card-text">{props.date}</p>
        <p className="open-sans card-text">{props.text}</p>
        <p className="open-sans card-text hidden-until-hover">{props.description}</p>
      </div>
    </div>
  )
}

Announcement.propTypes = {
  id: PropTypes.number,
  header: PropTypes.string,
  text: PropTypes.string,
  date: PropTypes.string,
  orgID: PropTypes.number
};

export default Announcement;