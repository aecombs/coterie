import React from 'react';
import SideNav from './SideNav';
import Members from './Members';
import Organization from './Organization';

const Dashboard = () => {
  return (
    <div>
      <p>Dashboard</p>
      <Organization />
      <SideNav />
      <Members />
    </div>
  )
}

export default Dashboard;