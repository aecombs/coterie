import React from 'react';
import Members from './Members';
import PropTypes from 'prop-types';

const Organization = (props) => {
  const formatDate = (date) => {
    const month = date.substring(5, 7)
    const day = date.substring(8, 10)
    const year = date.substring(0, 4)
    const string = `${year}-${month}-${day}`
    const formattedDate = new Date(string).toUTCString();
    return formattedDate.substring(0, 17)
  }

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

  return (
    <section className="w-100 container py-2">
      <div className="row">
        <div>
          <h4 className="">{props.name}</h4>
          <p className="open-sans">{props.missionStatement}</p>
        </div>
      </div>
      <div className="row justify-content-between">
        <p className="open-sans">Current Funds: {formatCurrency(props.totalFunds)}</p>
        <p className="open-sans">Established {formatDate(props.createdAt)}</p>
      </div>
      <div className="row">
        <Members 
        orgID={props.id}
        userID={props.userID}
        />
      </div>
    </section>
  )
}

Organization.propTypes = {
  id: PropTypes.number,
  name: PropTypes.string,
  missionStatement: PropTypes.string,
  totalFunds: PropTypes.number,
  user: PropTypes.number
};

export default Organization;