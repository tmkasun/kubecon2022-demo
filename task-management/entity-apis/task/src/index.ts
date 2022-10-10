import { PrismaClient } from '@prisma/client'
import express from 'express'
import dotenv from 'dotenv'


dotenv.config();
const prisma = new PrismaClient();
const app = express();

app.use(express.json());

app.get(`/task`, async (req, res) => {
  const { title } = req.query;

  const tasks = await prisma.task.findMany({
    where: {
      title: {
        contains: title as string
      }

    }
  })

  res.json(tasks);

})


app.get(`/task/:id`, async (req, res) => {
  const { id } = req.params;
  const task = await prisma.task.findUnique({
    where: {
      id: Number(id),
    },
  })
  if (!task) {
    res.sendStatus(404);
  } else {
    res.json(task);
  }
})

app.post(`/task`, async (req, res) => {
  const { title } = req.body;
  const task = await prisma.task.create({
    data: {
      title: title,
    }
  });
  res.json(task);
})

app.put(`/task/:id`, async (req, res) => {
  const { id } = req.params;
  const { title } = req.body;
  const task = await prisma.task.update({
    where: {
      id: Number(id)
    },
    data: {
      title: title
    }
  })
  res.json(task);
})

app.delete(`/task/:id`, async (req, res) => {
  const { id } = req.params
  const task = await prisma.task.delete({
    where: {
      id: Number(id),
    },
  })
  if (!task) {
    res.sendStatus(404);
  } else {
    res.json(task);
  }
})

const server = app.listen(8080, () =>
  console.log(`Server ready at: http://localhost:8080`),
)
