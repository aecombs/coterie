import React from 'react';
import PropTypes from 'prop-types';

const Member = (props) => {
  return (
    <div className="card">
      <div className="mx-2 card-body row justify-content-between">
        <p className="card-title">Name: {props.name}</p>
        <p className="card-title">Class: {props.class}</p>
      </div>
      <div className="mx-2 card-body row justify-content-between">
        <p className="card-text">${props.fundsRaised}</p>
        <p className="card-text">{props.birthdate}</p>
        <p className="card-text">{props.email}</p>
      </div>
    </div>
  )
}

Member.propTypes = {
  id: PropTypes.number,
  name: PropTypes.string,
  birthdate: PropTypes.string,
  class: PropTypes.string,
  email: PropTypes.string,
  fundsRaised: PropTypes.number,
  orgID: PropTypes.number
};

export default Member;