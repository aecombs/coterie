import React from 'react';
import PropTypes from 'prop-types';

const Member = (props) => {
  return (
    <tr className="text-center">
      <td><p className="card-title">{props.name}</p></td>
      <td><p className="card-title">{props.class}</p></td>
      <td><p className="card-text">${props.fundsRaised/100}</p></td>
      <td><p className="card-text">{props.email}</p></td>
      <td><p className="card-text">{props.birthdate}</p></td>
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