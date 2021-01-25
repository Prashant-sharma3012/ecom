import styled from 'styled-components'

const AppBar = styled.div`
    display: flex;
    min-height: 50px;
    background-color: ${props => props.bgcolor ? props.bgcolor : 'black'};
    color: ${props => props.color ? props.color : 'white'};
    position: fixed
    top: 0;
    left: 0;
    width: 100%;
`;

export default AppBar