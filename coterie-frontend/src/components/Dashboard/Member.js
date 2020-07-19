import React from 'react';
import PropTypes from 'prop-types';

const Member = (props) => {
  return (
    <tr className="text-center">
      <td className="card-title">{props.name}</td>
      <td className="card-title">{props.class}</td>
      <td className="card-text">${props.fundsRaised/100}</td>
      <td className="card-text">{props.email}</td>
      <td className="card-text">{props.birthdate}</td>
    </tr>
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