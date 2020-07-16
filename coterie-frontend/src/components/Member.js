import React from 'react';
import PropTypes from 'prop-types';

const Member = (props) => {
  return (
    <div className="list-group-item">
      <p className="">{props.name}{props.class}</p>
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