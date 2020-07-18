import React from 'react';
import PropTypes from 'prop-types';

const Chapter = (props) => {
  return (
    <div className="list-group-item card">
      <div className="card-body">
        <p className="open-sans card-title">{props.name}</p>
      </div>
      <div className="card-body">
        <p className="open-sans card-text hidden-until-hover">{props.text}</p>
      </div>
    </div>
  )
}

Chapter.propTypes = {
  id: PropTypes.number,
  name: PropTypes.string,
  text: PropTypes.string,
  position: PropTypes.string,
  scripID: PropTypes.number
};


export default Chapter;