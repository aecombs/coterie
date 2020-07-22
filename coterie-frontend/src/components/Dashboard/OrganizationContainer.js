import React, { useState, useEffect } from 'react';
import Organization from './Organization';
import OrgForm from './OrgForm';
import axios from 'axios';


const OrganizationContainer = (props) => {
  const [organizationList, setOrganizationList] = useState(null);
  const [errorMessage, setErrorMessage] = useState(null);
  const [addOrgMode, setAddOrgMode] = useState(false);

  const url = `${process.env.REACT_APP_API_BASE_URL}/users/${props.userID}/organizations/`
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

    //toggle visibility of new member component
    const toggleFormVisibility = () => {
      setAddOrgMode(!addOrgMode);
      return;
    }  

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
    />)
    })
  }


  return (
    <section className="d-flex container flex-wrap">
       <p className={ orgComponents !== undefined ? "hidden" : "open-sans" }>Organization</p>
       <button className={ orgComponents === undefined ? "btn btn-secondary float-right mb-2" : "hidden"} onClick={toggleFormVisibility}>+</button>
       <OrgForm 
       addOrgCallback={props.addOrgCallback}
       visibility={orgComponents === undefined}
       onSubmitCallback={toggleFormVisibility}
       />
      {orgComponents}
    </section>
  )
}

export default OrganizationContainer;