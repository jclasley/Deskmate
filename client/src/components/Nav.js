import React, {  useState  } from 'react';
import { 
    Collapse,
    Navbar,
    NavbarToggler,
    NavbarBrand,
    Nav,
    NavItem,
    NavLink,
    UncontrolledDropdown,
    DropdownToggle,
    DropdownMenu,
    DropdownItem,
    NavbarText
  } from 'reactstrap';
import {SlackConnect} from './SlackConnect.js';
import { ZendeskConnect } from './ZendeskConnect.js';

const Navigation = (props) => {
  const [isOpen, setIsOpen] = useState(false);

  const toggle = () => setIsOpen(!isOpen);

  return (
    <div>
      <Navbar color="light" light expand="md">
        <NavbarBrand href="/">deskmate</NavbarBrand>
        <NavbarToggler onClick={toggle} />
        <Collapse isOpen={isOpen} navbar>
          

          <NavbarText><SlackConnect /> <ZendeskConnect /></NavbarText>
        </Collapse>
      </Navbar>
    </div>
  );
}

export default Navigation;