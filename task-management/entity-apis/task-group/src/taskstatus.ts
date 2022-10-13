import express from 'express'
import prisma from './dbClient'
import { StatusCodes, getReasonPhrase } from 'http-status-codes';
import { Prisma, PrismaClient } from '@prisma/client';

const taskStatusRoute = express.Router();

taskStatusRoute.get(`/`, async (req, res) => {
  const { userId } = req.query;
  console.log("querying for user id : " + userId);

  const taskStatus = await prisma.taskStatus.findMany({
    where: {
      userId: userId as string
    },
    orderBy: {
      updatedAt: 'asc'
    }
  })

  res.status(StatusCodes.OK).json(taskStatus);
})


taskStatusRoute.get(`/:id`, async (req, res) => {
  const { id } = req.params;
  const taskStatus = await prisma.taskStatus.findUnique({
    where: {
      id: Number(id),
    },
  })
  if (!taskStatus) {
    res.status(StatusCodes.NOT_FOUND).send(getReasonPhrase(StatusCodes.NOT_FOUND));
  } else {
    res.status(StatusCodes.OK).json(taskStatus);
  }
})

taskStatusRoute.post(`/`, async (req, res) => {
  const { name, userId } = req.body;
  const taskStatus = await prisma.taskStatus.create({
    data: {
      name: name,
      userId: userId
    }
  });
  res.status(StatusCodes.CREATED).json(taskStatus);
})

taskStatusRoute.put(`/:id`, async (req, res) => {
  const { id } = req.params;
  const { name } = req.body;

  try {
    const taskStatus = await prisma.taskStatus.update({
      where: {
        id: Number(id)
      },
      data: {
        name: name
      }
    })

    res.status(StatusCodes.OK).json(taskStatus);

  } catch (e) {
    if (e instanceof Prisma.PrismaClientKnownRequestError) {
      console.log(e.message);
      res.status(StatusCodes.NOT_FOUND).send(getReasonPhrase(StatusCodes.NOT_FOUND));
    } else {
      throw e;
    }
  }
})

taskStatusRoute.delete(`/:id`, async (req, res) => {
  const { id } = req.params

  try {
    const taskStatus = await prisma.taskStatus.delete({
      where: {
        id: Number(id),
      },
    })
    res.status(StatusCodes.OK).json(taskStatus);

  } catch (e) {
    if (e instanceof Prisma.PrismaClientKnownRequestError) {
      console.log(e.message);
      res.status(StatusCodes.NOT_FOUND).send(getReasonPhrase(StatusCodes.NOT_FOUND));
    } else {
      throw e;
    }
  }

})

export { taskStatusRoute };   