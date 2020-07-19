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

  return (
    <div className="w-100">
      <div className="row justify-content-between">
        <div>
          <h4>{props.name}</h4>
          <p className="open-sans">{props.missionStatement}</p>
        </div>
        <div>
          <button onClick="" className="btn btn-secondary">+</button>
        </div>
      </div>
      <p className="open-sans">Established {formatDate(props.createdAt)}</p>
      <p className="open-sans">Current Funds: ${props.totalFunds/100}</p>
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