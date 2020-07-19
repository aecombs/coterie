import React, { useState, useEffect } from 'react';
import Organization from './Organization';
import OrgForm from './OrgForm';
import axios from 'axios';


const OrganizationContainer = (props) => {
  //TODO: logic that checks the URL... is it just the dashboard? then show xyz stuff.

  const [organizationList, setOrganizationList] = useState(null);
  const [errorMessage, setErrorMessage] = useState(null);

  const url = `http://localhost:3000/users/${props.userID}/organizations/`
  useEffect(() => {
    axios.get(url)
      .then( (response) => {
        const currentOrgs = response.data;
        setOrganizationList(currentOrgs);
      })
      .catch((err) => {
        console.log(`There was an error: ${err["message"]}`)
      });
  },[url])

  // const addOrg = (org) => {
  //   axios.post(url, org)
  //   .then((response) => {
  //     setErrorMessage(`Organization ${org} added`);
  //   })
    
  //   .catch((error) => {
  //     setErrorMessage(error.message);
  //     console.log(`Unable to add org: ${errorMessage}`);
  //   })
  // }


  let orgComponents = undefined
  if (organizationList !== null && organizationList.length > 0) {
    orgComponents = organizationList.map((org) => {
    return(<Organization
      key = {org.id}
      id = {org.id}
      name = {org.name}
      missionStatement = {org.mission_statement}
      totalFunds = {org.total_funds}
      createdAt = {org.created_at}
      userID = "1"
    />) 
    })
  }


  return (
    <section className="d-flex container flex-wrap">
       <p className={ orgComponents !== undefined ? "hidden" : "open-sans" }>It looks like you don't have an organization yet. Go ahead and make a new one!</p>
       <OrgForm addOrgCallback={props.addOrgCallback} />
      {orgComponents}
    </section>
  )
}

export default OrganizationContainer;