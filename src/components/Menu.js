import React from "react";
import { Menu, Layout } from "antd";

export default class CMenu extends React.Component {
  handleMenuClick = e => {
    window.location.href = e.key;
  };
  render() {
    const { Item } = Menu;
    const { Sider, Content, Header } = Layout;
    const {Component} = this.props;
    return (
      <Layout>
        <Sider>
          <Menu  onClick={this.handleMenuClick}>
            <Item key="/company">
              <span>Company</span>
            </Item>
            <Item key="/position">
              <span>Position</span>
            </Item>
            <Item key="/candidate">
              <span>Candidate</span>
            </Item>
            <Item key="/task">
              <span>Task</span>
            </Item>
            <Item key="/folder">
              <span>Folder</span>
            </Item>
          </Menu>
        </Sider>
        <Layout>
          <Content><Component /></Content>
        </Layout>
      </Layout>
    );
  }
}
