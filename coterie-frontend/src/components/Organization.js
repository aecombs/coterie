import React from 'react';
import Members from './Members';
import PropTypes from 'prop-types';

const Organization = (props) => {

  return (
    <div className="w-100">
      <h2 className="open-sans">{props.name}</h2>
      <p className="open-sans">{props.missionStatement}</p>
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