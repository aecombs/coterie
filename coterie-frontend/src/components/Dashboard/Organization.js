import React, { useState } from 'react';
import Members from './Members';
import PropTypes from 'prop-types';
import OrgForm from './OrgForm';

const Organization = (props) => {
  const [updateMode, setUpdateMode] = useState(false);

  const updateOrg = (orgObj) => {
    props.submitOrgCallback(orgObj)
  }

  const setUpdate = () => {
    setUpdateMode(!updateMode);
  }



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
    <section className="w-100 container py-2 card">
      <div className="row card-body">
        <div className="w-100">
          <button onClick={setUpdate} className={ updateMode ? "hidden" : "btn list-group-item list-group-item-action"}>
          <h4 className="">{props.name}</h4>
          <p className="open-sans">{props.missionStatement}</p>
          </button>
          <OrgForm 
           id={props.id}
           name={props.name}
           missionStatement={props.missionStatement}
           userID={props.userID}
           visibility={updateMode}
           submitOrgCallback={updateOrg}
           onSubmitCallback={setUpdate}
          />
        </div>
      </div>
      <div className="row justify-content-between card-body">
        <p className="open-sans">Current Funds: {formatCurrency(props.totalFunds)}</p>
        <p className="open-sans">Established {formatDate(props.createdAt)}</p>
      </div>
      <div className="row card-body">
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