import Table from 'react-bootstrap/Table'
import Button from 'react-bootstrap/Button';
import Modal from 'react-bootstrap/Modal'
import React from 'react'
import { useCalorieStore } from '../utils/useStore';



function StripedTable(): JSX.Element{
  const {calorieData, totalCalories, fetchCalorieData, fetchTotalCalories} = useCalorieStore()
  const [food, setFood] = React.useState<string | null>('');
  const [calorie, setCalorie] = React.useState<number | null>(null);

  // React Modal

  const [show, setShow] = React.useState<boolean>(false);

  // React form functions

  const handleShow = () => setShow(true);
  const handleClose = () => setShow(false);

  const handleUpdate = (id: number, e: React.FormEvent) => {
    e.preventDefault()
    setShow(false)
    if (food !== '' && calorie !== null) {
      fetch(`${import.meta.env.VITE_PORT}/calories/${id}`,
      {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          food: food,
          calorie:calorie
        })
        })
        .then(res => res.json())
        .then(() => {
          fetchCalorieData()
          fetchTotalCalories()
        })
    }
  }
  
  const handleDelete = (id: number) => {
    fetch(`${import.meta.env.VITE_PORT}/calories/${id}`,
      {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json'
        }
      })
      .then(res => res.json())
      .then(() => {
        fetchCalorieData()
        fetchTotalCalories()
      })
  }
  return (
    <Table striped bordered hover>
      <thead>
        <tr>
          <th>FOOD</th>
          <th>CALORIES</th>
          <th colSpan={2} className='endRow'>Make Changes</th>
        </tr>
      </thead>
      <tbody>
        {
          calorieData?.map((data, index) => {
            return (
              <tr key={index}>
                <td>{data.food}</td>
                <td>{data.calorie}</td>
                <td>
                  <Button
                    variant="primary"
                    onClick={handleShow}>
                    Edit
                  </Button>
                </td>
                <td>
                  <Button
                    variant="primary"
                    onClick={() => {
                      handleDelete(data.id);
                    }}
                  >
                    DELETE
                  </Button>
                </td>
                <Modal show={show} onHide={handleClose} animation={false}>
                  <Modal.Header closeButton>
                    <Modal.Title>Update today's diet</Modal.Title>
                  </Modal.Header>
                  <form onSubmit={(e:React.FormEvent) => {
                    handleUpdate(data.id, e)
                  }}>
                    <Modal.Body>
                      <div className="form-group">
                        <label htmlFor="food">Food</label>
                        <input
                          type="text"
                          className="form-control"
                          id="food"
                          placeholder={data.food}
                          value={food?.toString()}
                          onChange={(
                            e: React.ChangeEvent<HTMLInputElement>
                          ) => {
                            setFood(e.target.value);
                          }}
                        />
                        <label htmlFor="calorie">Calorie</label>
                        <input
                          type="number"
                          min="0"
                          className="form-control"
                          id="calorie"
                          placeholder={(data.calorie).toString()}
                          value={calorie?.toString()}
                          onChange={(
                            e: React.ChangeEvent<HTMLInputElement>
                          ) => {
                            setCalorie(parseInt(e.target.value));
                          }}
                        />
                      </div>
                    </Modal.Body>
                    <Modal.Footer>
                      <Button variant="primary" type="submit"
                        disabled={food === '' || calorie === null}
                      >
                        Save Changes
                      </Button>
                    </Modal.Footer>
                  </form>
                </Modal>
              </tr>
            );
          })
        }
        <tr>
          <td colSpan={3} className='endRow'>Total Calories</td>
          <td colSpan={2}>{totalCalories}</td>
        </tr>
      </tbody>
    </Table>
  );
}

export default StripedTable