import React from 'react';
import Members from './Members';
import PropTypes from 'prop-types';

const Organization = (props) => {
  const formatDate = (date) => {
    const month = date.substring(5, 7)
    const day = date.substring(8, 10)
    const year = date.substring(0, 4)
    const formatedDate = `${month}-${day}-${year}`
    return formatedDate
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
    <div className="w-100">
      <div className="row justify-content-between">
        <div>
          <h4>{props.name}</h4>
          <p className="open-sans">{props.missionStatement}</p>
        </div>
        {/* <div>
          <a href="/dashboard/new" className="">
          <button className="btn btn-secondary">+</button>
          </a>
        </div> */}
      </div>
      <p className="open-sans">Established {formatDate(props.createdAt)}</p>
      <p className="open-sans">Current Funds: {formatCurrency(props.totalFunds)}</p>
      <Members 
      orgID={props.id}
      userID={props.userID}
      />
    </div>
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