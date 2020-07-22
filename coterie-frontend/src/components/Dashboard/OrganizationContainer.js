import React, { useState, useEffect } from 'react';
import Organization from './Organization';
import OrgForm from './OrgForm';
import axios from 'axios';


const OrganizationContainer = (props) => {
  const [organizationList, setOrganizationList] = useState(null);
  const [errorMessage, setErrorMessage] = useState(null);
  const [addOrgMode, setAddOrgMode] = useState(false);


  const url = `${process.env.REACT_APP_API_BASE_URL}/users/${props.userID}/organizations`

  const addOrg = (orgObj) => {
    delete orgObj["id"];

    axios.post(url, orgObj)
    .then((response) => {
      setErrorMessage(`organization ${orgObj["name"]} added`);
      //TODO: Update to use intended redirect URL from response
      window.location.reload();
    })
    
    .catch((error) => {
      setErrorMessage(error.message);
      console.log(`Unable to add org: ${errorMessage}`);
    })
  }


  const updateOrg = (orgObj) => {
    axios.put(`${url}/${orgObj.id}`, orgObj)
    .then((response) => {
      setErrorMessage(`Organization ${orgObj["name"]} was updated`);
      window.location.reload();
    })
    
    .catch((error) => {
      setErrorMessage(error.message);
      console.log(`Unable to update organization: ${errorMessage}`);
    })
  }

  //toggle visibility of new member component
  const toggleFormVisibility = () => {
    setAddOrgMode(!addOrgMode);
    return;
  }  


  useEffect(() => {
    axios.get(url)
      .then( (response) => {
        const currentOrgs = response.data;
        setOrganizationList(currentOrgs);
      })
      .catch((err) => {
        setErrorMessage(err);
        console.log(`There was an error: ${err["message"]}`)
      });
  },[url])

  let orgComponents = undefined
  if (organizationList !== null && organizationList.length > 0) {
    orgComponents = organizationList.map((org) => {
    props.tellParentID(org.id);
    return(<Organization
      key = {org.id}
      id = {org.id}
      name = {org.name}
      missionStatement = {org.mission_statement}
      totalFunds = {org.total_funds}
      createdAt = {org.created_at}
      userID = {props.userID}
      submitOrgCallback={updateOrg}
    />)
    })
  }


  return (
    <section className="d-flex container flex-wrap">
      <div className="d-flex py-2 justify-content-between w-100">
        <h4 className="font-weight-bolder">
          Organization</h4>
        <button className={ orgComponents === undefined ? "btn btn-outline-secondary mb-2" : "hidden"} onClick={toggleFormVisibility}>+</button>
      </div>

      <OrgForm 
        userID={props.userID}
        addOrgCallback={props.addOrgCallback}
        visibility={addOrgMode}
        onSubmitCallback={toggleFormVisibility}
        submitOrgCallback={addOrg}
       />
      {orgComponents}
    </section>
  )
}

export default OrganizationContainer;