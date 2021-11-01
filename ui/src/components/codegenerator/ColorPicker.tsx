import React from 'react';
import reactCSS from 'reactcss';
import { ChromePicker } from 'react-color';

const ColorPicker = ({ ...props }) => {
  const [state, setState] = React.useState({
    displayColorPicker: false
  });

  const handleClick = () => {
    setState({ displayColorPicker: !state.displayColorPicker });
  };

  const handleClose = () => {
    setState({ displayColorPicker: false });
  };

  const styles = reactCSS({
    default: {
      color: {
        width: '36px',
        height: '14px',
        borderRadius: '2px',
        background: `${props.color}`
      },
      swatch: {
        padding: '5px',
        background: '#fff',
        borderRadius: '1px',
        boxShadow: '0 0 0 1px rgba(0,0,0,.1)',
        display: 'inline-block',
        cursor: 'pointer'
      },
      popover: {
        position: 'absolute',
        zIndex: '2'
      },
      cover: {
        position: 'fixed',
        top: '0px',
        right: '0px',
        bottom: '0px',
        left: '0px'
      }
    }
  });

  return (
    <div>
      <div style={styles.swatch} onClick={handleClick}>
        <div style={styles.color} />
      </div>
      {state.displayColorPicker ? (
        <div style={styles.popover}>
          <div style={styles.cover} onClick={handleClose} />
          <ChromePicker color={props.color} onChange={props.handleChange} />
        </div>
      ) : null}
    </div>
  );
};

export default ColorPicker;
