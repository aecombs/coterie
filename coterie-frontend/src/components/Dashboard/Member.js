import React from 'react';
import PropTypes from 'prop-types';

const Member = (props) => {
  const formatCurrency = (int) => {
    let modInt = 0;
    if (int !== 0) {
      modInt = int/100
    }
    if (int === undefined || int === null) {
      modInt = 0;
    }
    const formattedInt = new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' }).format(modInt)
    return formattedInt
  }

  const formatDate = (date) => {
    const month = date.substring(5, 7)
    const day = date.substring(8, 10)
    const year = date.substring(0, 4)
    const string = `${year}-${month}-${day}`
    const formattedDate = new Date(string).toUTCString();
    return formattedDate.substring(0, 17)
  }

  return (
    <tr className="text-left">
      <td className="card-title open-sans">{props.name}</td>
      <td className="card-title open-sans">{props.class}</td>
      <td className="card-text open-sans">{formatCurrency(props.fundsRaised)}</td>
      <td className="card-text open-sans">{props.email}</td>
      <td className="card-text open-sans">{formatDate(props.birthdate)}</td>
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