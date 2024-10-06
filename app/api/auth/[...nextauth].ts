import NextAuth from "next-auth";
import GoogleProvider from "next-auth/providers/google";

export const authOptions = {
  providers: [
    GoogleProvider({
      clientId: process.env.GOOGLE_CLIENT_ID!,
      clientSecret: process.env.GOOGLE_CLIENT_SECRET!,
    }),
  ],
  callbacks: {
    async signIn({ user, account, profile }) {
      // Here you can add logic to check if the user is a super user or needs approval
      return true;
    },
    async session({ session, user }) {
      // Add custom session properties here
      return session;
    },
  },
};

export default NextAuth(authOptions);