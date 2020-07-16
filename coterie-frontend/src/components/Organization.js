import React from 'react';
import Members from './Members';
import PropTypes from 'prop-types';

const Organization = (props) => {
  return (
    <div className="">
      <h2 className="open-sans">{props.name}</h2>
      <p className="open-sans">{props.mission_statement}</p>
      {/* <Collection
      selection="member" /> */}
      <Members 
      org_id={props.id}/>
    </div>
  )
}

Organization.propTypes = {
  id: PropTypes.number,
  name: PropTypes.string,
  mission_statement: PropTypes.string,
  total_funds: PropTypes.number,
  user_id: PropTypes.number
};

export default Organization;